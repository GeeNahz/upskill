defmodule MiniRedis do
    use GenServer

  def init(_) do
    {:ok, %{}}
  end

  def start_link(opts \\ []) do
    GenServer.start_link(__MODULE__, [], opts)
  end

  def set(key, value) do
    GenServer.call(__MODULE__, {:set, key, value})
  end

  def get(key) do
    GenServer.call(__MODULE__, {:get, key})
  end

  def handle_call({:set, key, value}, _from, state) do
    {:reply, :ok, Map.merge(state, %{key => value})}
  end

  def handle_call({:get, key}, _from, state) do
    {:reply, Map.fetch(state, key), state}
  end
end
