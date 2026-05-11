defmodule Math do
  @spec sum(integer, integer) :: integer
  def sum(a, b), do: a + b

  @spec div(integer, integer) :: {:ok, integer} | {:error, String}
  def div(a, b) do
    # ...
  end
end

defmodule Customer do
  @type entity_id() :: integer()

  @type t :: %Customer{id: entity_id(), first_name: String.t(), last_name: String.t()}

  # full form
  # defstruct [{:id, 0}, {:first_name, nil}, {:last_name, nil}]

  # short and common form
  defstruct id: 0, first_name: nil, last_name: nil, age: 0
end

defmodule CustomerDao do
  @type reason :: String.t()

  @spec get_customer(Customer.entity_id()) :: {:ok, Customer} | {:error, reason}

  def get_customer(id) do
    # ...
    IO.puts("GETTING CUSTOMER #{id}")
  end
end

defmodule Person do
  defstruct first_name: nil, last_name: nil, age: 0
end
