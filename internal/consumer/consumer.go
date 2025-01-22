package consumer

import (
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
	"github.com/goccy/go-json"
	"log/slog"

	"github.com/IBM/sarama"
)

type RecipesConsumer struct {
	Ready        chan bool
	recipesSaver RecipesSaver
	logger       *slog.Logger
}

func NewRecipesConsumer(saver RecipesSaver, log *slog.Logger) *RecipesConsumer {
	return &RecipesConsumer{Ready: make(chan bool), recipesSaver: saver, logger: log}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *RecipesConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *RecipesConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *RecipesConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				consumer.logger.Info("message channel was closed")
				return nil
			}
			consumer.logger.Info("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			recipe := entity.CreateRecipe{}
			recipe.Ingredients = make([]entity.Ingredient, 0)
			err := json.Unmarshal(message.Value, recipe)
			if err != nil {
				consumer.logger.Error("cannot unmarshal massage", slog.Any("error :", err))
			}

			if err := consumer.recipesSaver.SaveRecipe(session.Context(), recipe); err != nil {
				return err
			}

			session.MarkMessage(message, "")
			session.Commit()
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
