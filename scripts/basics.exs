defmodule MyFunction do
  # functions here
  def sum(a, b) do
    a + b
  end

  def sub(a, b) do
    a - b
  end

  def print_result(f) do
    IO.puts(f.())
  end
end

defmodule DoSomeMath do
  import MyFunction, only: [sum: 2, sub: 2]

  def add_and_subtract(a, b, c) do
    sub(sum(a, b), c)
  end
end

defmodule Bot do
  def greet("") do
    IO.puts("None to greet.")
  end

  def greet(name) do
    IO.puts("Hell0 #{name}")
  end
end

defmodule Factorial do
  # Tail-call optimisation
  def do_it(n) do
    internal_do(n, 1)
  end

  defp internal_do(0, acc), do: acc

  defp internal_do(n, acc) do
    internal_do(n - 1, acc * n)
  end
end

defmodule ListUtils do
  def sum([]), do: 0

  def sum([h | t]) do
    h + sum(t)
  end
end

defmodule Foo do
  def divide_by_10(value) when value > 0 and (is_float(value) or is_integer(value)) do
    value / 10
  end
end
