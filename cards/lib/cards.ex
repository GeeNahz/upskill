# defmodule Cards do
#   @moduledoc """
#   Documentation for `Cards`.
#   """
#
#   @doc """
#   Hello world.
#
#   ## Examples
#
#       iex> Cards.hello()
#       :world
#
#   """
#   def hello do
#     :world
#   end
# end

defmodule Cards do
  @moduledoc """
  Provides methods for creating and handlign a deck of cards
  """

  @doc """
  Returns a list of strings representing a deck of playing cards
  """
  def create_deck() do
    values = [
      "Ace",
      "Two",
      "Three",
      "Four",
      "Five",
      "Six",
      "Seven",
      "Eight",
      "Nine",
      "Ten",
      "Joker",
      "Queen",
      "King"
    ]

    suits = ["Hearts", "Spades", "Diamonds", "Clubs"]

    for value <- values, suit <- suits do
      "#{value} of #{suit}"
    end
  end

  @doc """
  Determines whether a deck contains a given card

  ## Examples
      
      iex> deck = Cards.create_deck()
      iex> Cards.contains?(deck, "Ace of Hearts")
      iex> true
  """
  def contains?(deck, card) do
    Enum.member?(deck, card)
  end

  def shuffle(deck) do
    Enum.shuffle(deck)
  end

  @doc """
  Divides a deck into a hand and the remainder of the deck.
  The `hand_size` argument indicates how many cards should be in the hand

  ## Examples

      iex> deck = Cards.create_deck()
      iex> {hand, deck} = Cards.deal(deck, 4)
      iex> hand
      ["Ace of Hearts", "Ace of Spades", "Ace of Diamonds", "Ace of Clubs"]

  """
  def deal(deck, hand_size) do
    Enum.split(deck, hand_size)
  end

  def save(deck, filename) do
    binary = :erlang.term_to_binary(deck)
    File.write(filename, binary)
  end

  def load(filename) do
    case File.read(filename) do
      {:ok, binary} -> :erlang.binary_to_term(binary)
      {:error, :enoent} -> "That file does not exist."
      {:error, reason} -> "Something went wrong: #{reason}"
    end
  end

  def create_hand(hand_size) do
    Cards.create_deck()
    |> Cards.shuffle()
    |> Cards.deal(hand_size)
  end
end
