defmodule HelloWeb.PageController do
  use HelloWeb, :controller

  def home(conn, _params) do
    # render(conn, :home)
    conn
    |> put_flash(:error, "Let's pretend we have an error")
    |> redirect(to: ~p"/redirect_test")

    # |> render(:home)

    # redirect(conn, to: ~p"/redirect_test") # internal redirect with 'to' 
    # redirect(conn, external: "https://elixir-lang.org/") # internal redirect with 'external'
  end

  def redirect_test(conn, _params) do
    render(conn, :home)
  end
end
