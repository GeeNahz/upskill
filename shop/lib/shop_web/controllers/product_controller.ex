defmodule ShopWeb.ProductController do
  use ShopWeb, :controller

  def index(conn, _params) do
    # conn
    # |> redirect(to: ~p"/products")
    render(conn, :index)
  end

  def show(conn, %{"id" => id}) do
    render(conn, :show, id: id)
  end
end
