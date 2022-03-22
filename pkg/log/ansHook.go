package log

import (
	"fmt"
	"github.com/SAP/jenkins-library/pkg/ans"
	"github.com/SAP/jenkins-library/pkg/xsuaa"
	"github.com/sirupsen/logrus"
)

// ANSHook is used to set the hook features for the logrus hook
type ANSHook struct {
	correlationID string
	client        ans.Client
	event         ans.Event
}

// NewANSHook creates a new ANS hook for logrus
func NewANSHook(serviceKey, correlationID, eventTemplate string) ANSHook {
	ansServiceKey, err := ans.UnmarshallServiceKeyJSON(serviceKey)
	if err != nil {
		Entry().Warnf("cannot initialize ans due to faulty serviceKey json: %v", err)
	}
	var event ans.Event
	if len(eventTemplate) > 0 {
		event, err = ans.UnmarshallEventJSON(eventTemplate)
		if err != nil {
			Entry().Warnf("provided ANS event template could not be unmarshalled: %v", err)
		}
	}
	x := xsuaa.XSUAA{
		OAuthURL:     ansServiceKey.OauthUrl,
		ClientID:     ansServiceKey.ClientId,
		ClientSecret: ansServiceKey.ClientSecret,
	}
	h := ANSHook{
		correlationID: correlationID,
		client:        ans.ANS{XSUAA: x, URL: ansServiceKey.Url},
		event:         event,
	}
	return h
}

// Levels returns the supported log level of the hook.
func (ansHook *ANSHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.DebugLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel}
}

// Fire creates a new event from the logrus and sends an event to the ANS backend
func (ansHook *ANSHook) Fire(entry *logrus.Entry) error {
	if len(ansHook.event.EventType) == 0 {
		ansHook.event.EventType = "Piper"
	}
	ansHook.event.EventTimestamp = entry.Time.Unix()
	ansHook.event.Severity, ansHook.event.Category = ans.TranslateLogrusLogLevel(entry.Level)
	if ansHook.event.Subject == "" {
		ansHook.event.Subject = fmt.Sprint(entry.Data["stepName"])
	}
	ansHook.event.Body = entry.Message
	if len(ansHook.event.Tags) == 0 {
		ansHook.event.Tags = make(map[string]interface{})
	}
	ansHook.event.Tags["logLevel"] = entry.Level.String()
	ansHook.event.Tags["ans:correlationId"] = ansHook.correlationID
	for k, v := range entry.Data {
		ansHook.event.Tags[k] = v
	}

	err := ansHook.client.Send(ansHook.event)
	if err != nil {
		return err
	}
	return nil
}