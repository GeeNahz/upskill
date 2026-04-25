defmodule HelloWeb.Plug.Locale do
  import Plug.Conn

  @locales ["en", "fr", "de"]

  def init(default), do: default

  def call(%Plug.Conn{params: %{"locale" => loc}} = conn, _default) when loc in @locales do
    conn
    |> assign(:locale, loc)
  end

  def call(conn, default) do
    conn
    |> assign(:locale, default)
  end
end
