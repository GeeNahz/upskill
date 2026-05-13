# Macros
# Abstract Syntax Trees (AST)

defmodule Logger do
  defmacro log(msg) do
    if is_log_enabled?() do
      quote do
        IO.puts("> From log: #{unquote(msg)}")
      end
    end
  end

  defp is_log_enabled?(), do: true
end
