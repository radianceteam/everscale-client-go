package client

// DebotFetch Downloads debot smart contract (code and data) from blockchain and creates
// an instance of Debot Engine for it.
// Remarks: It does not switch debot to context 0. Browser Callbacks are not called.
func (c *Client) DebotFetch(p *ParamsOfFetch) (*RegisteredDebot, error) {
	response := new(RegisteredDebot)
	err := c.dllClient.waitErrorOrResultUnmarshal("debot.fetch", p, response)

	return response, err
}

// DebotStart Downloads debot smart contract from blockchain and switches it to context zero.
// Returns a debot handle which can be used later in `execute` function.
// This function must be used by Debot Browser to start a dialog with debot.
// While the function is executing, several Browser Callbacks can be called,
// since the debot tries to display all actions from the context 0 to the user.
// Remarks: `start` is equivalent to `fetch` + switch to context 0.
func (c *Client) DebotStart(p *ParamsOfStart) (*RegisteredDebot, error) {
	response := new(RegisteredDebot)
	err := c.dllClient.waitErrorOrResultUnmarshal("debot.start", p, response)

	return response, err
}
