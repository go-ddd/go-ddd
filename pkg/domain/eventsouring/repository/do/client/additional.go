func (c *EventClient) Debug() *EventClient {
if c.debug {
return c
}
cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
return &EventClient{config: cfg}
}
