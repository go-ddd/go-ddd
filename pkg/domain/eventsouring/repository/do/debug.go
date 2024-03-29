// Code generated by ent, DO NOT EDIT.

package do

import "entgo.io/ent/dialect"

// Debug is enable debug mode.
func (c *EventClient) Debug() *EventClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), debug: true, log: c.log, hooks: c.hooks}
	return &EventClient{config: cfg}
}
