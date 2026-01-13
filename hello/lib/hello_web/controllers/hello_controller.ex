defmodule HelloWeb.HelloController do
  use HelloWeb, :controller

  def index(conn, _params) do
    render(conn, :index)
  end

  def show(conn, %{"messenger" => messenger}) do
    # can use `assign/3` to pass in params and pipe down to render
    conn
    |> assign(:messenger, messenger)
    |> assign(:receiver, "Dweezil")
    |> render(:show)

    # render(conn, :show, messenger: messenger, receiver: "Dweezil") # for html/template based response
    # json(conn, %{id: messenger}) # for json-base response
    # text(conn, "From messenger: #{messenger}") # for text-based response
  end
end
