
# Day 042 - Chapter 14 Exercise 3

**Chapter:** The Context

## Original Prompt

3. Assume you have a simple logging function that looks like this: func Log(ctx context.Context, level Level, message string) { var inLevel Level // TODO get a logging level out of the context and assign it to inLevel if level == Debug && inLevel == Debug { fmt.Println(message) } if level == Info && (inLevel == Debug || inLevel == Info) { fmt.Println(message) } } Define a type called Level whose underlying type is string. Define two con‐ stants of this type, Debug and Info, and set them to "debug" and "info", respectively. Create functions to store the log level in the context and to extract it. Create a middleware function to get the logging level from a query parameter called log_level. The valid values for log_level are debug and info. Finally, fill in the TODO in Log to properly extract the log level from the context. If the log level is not assigned or is not a valid value, nothing should be printed.

## Acceptance

Implement the logging helpers inside the local `log/` package. The acceptance test imports `log.ContextWithLevel`, `log.LevelFromContext`, `log.Log`, and `log.Middleware` directly.

Run `go test` inside this folder when you want feedback.
