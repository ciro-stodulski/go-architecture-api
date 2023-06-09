package verifyconsumer

import (
	domaindto "go-clean-api/cmd/domain/dto"
	usecase "go-clean-api/cmd/domain/use-case"
	consumer "go-clean-api/cmd/presentation/amqp/consumers"
	ports_amqp "go-clean-api/cmd/presentation/amqp/ports"
	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	verifiyConsumer struct {
		NotifyUserUseCase usecase.NotifyUserUseCase
	}
)

func New(NotifyUserUseCase usecase.NotifyUserUseCase) consumer.Comsumer {
	return &verifiyConsumer{
		NotifyUserUseCase,
	}
}

func (createConsumer *verifiyConsumer) GetConfig() consumer.ConsumeConfig {
	return consumer.ConsumeConfig{
		Queue:  "notify.create",
		Schema: domaindto.Event{},
	}
}

func (createConsumer *verifiyConsumer) MessageHandler(msg ports_amqp.Message) error {
	dto := domaindto.Event{}

	mapstructure.Decode(msg.Body, &dto)

	err := createConsumer.NotifyUserUseCase.Notify(dto)

	if err != nil {
		return err
	}

	return nil
}

func (createConsumer *verifiyConsumer) OnConsumerError(err error) consumer.AckConfig {
	log.Default().Println("error verify consumer:", err)

	return consumer.AckConfig{
		Multiple: false,
		Requeue:  false,
	}
}
