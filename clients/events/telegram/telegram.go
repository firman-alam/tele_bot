package telegram

import (
	"errors"

	"github.com/firman-alam/tele_bot.git/clients/events"
	"github.com/firman-alam/tele_bot.git/clients/telegram"
	"github.com/firman-alam/tele_bot.git/lib/e"
	"github.com/firman-alam/tele_bot.git/storage"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

type Meta struct {
	ChatID int
	Username string
}

var ErrUnknownEventType = errors.New("unknown event type")

func New(client *telegram.Client, storage storage.Storage) *Processor { 
	return &Processor{
		tg: nil,
		storage: nil,
	}	
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1 

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		p.processMessage(event)
	default:
		return e.WrapIfErr("can't process message", ErrUnknownEventType)
	}

	return nil
}

func (p *Processor) processMessage(event events.Event) {
	meta, err := meta(vent)
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID: upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	// chatID username
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}