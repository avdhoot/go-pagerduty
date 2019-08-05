package pagerduty

import (
	"fmt"
	"net/http"
)

// ListEventRuleResponse struct
type ListEventRuleResponse struct {
	Rules         []EventRule `json:"rules"`
	ObjectVersion string      `json:"object_version"`
	ID            string      `json:"id"`
	FormatVersion string      `json:"format_version"`
	ExternalID    string      `json:"external_id"`
}

// EventRule struct
type EventRule struct {
	ID                string        `json:"id,omitempty"`
	Disabled          bool          `json:"disabled,omitempty"`
	Condition         []interface{} `json:"condition,omitempty"`
	CatchAll          bool          `json:"catch_all,omitempty"`
	AdvancedCondition []interface{} `json:"advanced_condition,omitempty"`
	Actions           []interface{} `json:"actions,omitempty"`
}

// ListEventRule lists existing event rules.
func (c *Client) ListEventRule() (*ListEventRuleResponse, error) {
	resp, err := c.get("/event_rules")
	if err != nil {
		return nil, err
	}
	var result ListEventRuleResponse
	return &result, c.decodeJSON(resp, &result)
}

// CreateEventRule creates a new event rule.
func (c *Client) CreateEventRule(e EventRule) (*EventRule, error) {
	resp, err := c.post("/event_rules", e, nil)
	return getEventRuleFromResponse(c, resp, err)
}

// UpdateEventRule updates an existing event rule.
func (c *Client) UpdateEventRule(e EventRule) (*EventRule, error) {
	resp, err := c.put("/event_rules/"+e.ID, e, nil)
	return getEventRuleFromResponse(c, resp, err)
}

// DeleteEventRule deletes an existing event rule.
func (c *Client) DeleteEventRule(id string) error {
	_, err := c.delete("/event_rules//" + id)
	return err
}

func getEventRuleFromResponse(c *Client, resp *http.Response, err error) (*EventRule, error) {
	if err != nil {
		return nil, err
	}
	var target EventRule
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target, nil
}
