# Sentimental

This is a first pass at a Go CLI command that starts an OSC server and waits for a message that looks like `/text STRING`

When it receives that message, it does sentiment analysis on the string and sends the resulting float32 (negative number
is negative sentiment, positive number is positive sentiment) via OSC to a specified server address and port.
