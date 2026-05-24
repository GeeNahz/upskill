defmodule AsyncMath do
  def start() do
    IO.puts("Process started. Waiting to receive message...")

    receive do
      {:sum, [a, b], pid} -> send(pid, {:result, a + b})
    end

    start()
  end

  def add(a, b) do
    IO.puts("Sum: #{a + b}")
  end
end
