# Protocols
defprotocol Printable do
  def to_csv(data)
end

# Implementation
defimpl Printable, for: Map do
  def to_csv(map) do
    Map.keys(map)
    |> Enum.map(fn k -> map[k] end)
    |> Enum.join(",")
  end
end

defimpl Printable, for: List do
  def to_csv(list) do
    Enum.map(list, fn item -> Printable.to_csv(item) end)
  end
end

defimpl Printable, for: Integer do
  def to_csv(i) do
    to_string(i)
  end
end

# Behaviours
defmodule TalkingAnimal do
  @callback say(what :: String.t()) :: {:ok}
  @callback legs() :: {:ok}
end

defmodule Cat do
  @behaviour TalkingAnimal

  def say(str) do
    "miaooo"
  end

  def legs() do
    4
  end
end

defmodule Dog do
  @behaviour TalkingAnimal

  def say(str) do
    "woff"
  end

  def legs() do
    4
  end
end

defmodule Factory do
  def get_animal() do
    # can get module from configuration file
    Cat
  end
end
