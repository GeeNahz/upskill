defmodule Identicon do
  alias Vix.Vips.Image, as: VImage

  def main(input) do
    input
    |> has_input()
    |> pick_color()
    |> build_grid()
    |> filter_odd_squares()
    |> build_pixel_map()
    |> draw_image()
    |> save_image(input)
  end

  def save_image(svg, input) do
    # File.write("#{input}.png", image)
    {:ok, image} = VImage.new_from_buffer(svg, "", scale: 1)
    png = VImage.write_to_buffer(image, ".png")
    File.write("#{input}.png", png)
  end

  def draw_image(%Identicon.Image{color: {r, g, b}, pixel_map: pixel_map}) do
    color = "rgb(#{r}, #{g}, #{b})"

    rects =
      Enum.map(pixel_map, fn {{x1, y1}, {x2, y2}} ->
        width = x2 - x1
        height = y2 - y1

        rect = """
        <rect x="#{x1}" y="#{y1}" width="#{width}" height="#{height}" fill="#{color}" />
        """

        rect
      end)
      |> Enum.join("\n")

    svg = """
      <svg
      xlms="https://www.w3.org/2000/svg"
      height="200"
      width="200"
      viewBox="0 0 250 250"
      >
      <rect width="250" height="250" fill="white" />
      <g>
      #{rects}
      </g>
      </svg>
    """

    svg
  end

  def build_pixel_map(%Identicon.Image{grid: grid} = image) do
    pixel_map =
      Enum.map(grid, fn {_code, index} ->
        horizontal = rem(index, 5) * 50
        vertical = div(index, 5) * 50

        top_left = {horizontal, vertical}
        bottom_right = {horizontal + 50, vertical + 50}

        {top_left, bottom_right}
      end)

    %Identicon.Image{image | pixel_map: pixel_map}
  end

  def filter_odd_squares(%Identicon.Image{grid: grid} = image) do
    grid = Enum.filter(grid, fn {code, _index} -> rem(code, 2) == 0 end)
    %Identicon.Image{image | grid: grid}
  end

  def build_grid(%Identicon.Image{hex: hex} = image) do
    grid =
      hex
      |> Enum.chunk(3)
      |> Enum.map(&mirror_row/1)
      |> List.flatten()
      |> Enum.with_index()

    %Identicon.Image{image | grid: grid}
  end

  def mirror_row([first, second | _tail] = row) do
    row ++ [second, first]
  end

  def pick_color(%Identicon.Image{hex: [r, g, b | _tail]} = image) do
    %Identicon.Image{image | color: {r, g, b}}
  end

  def has_input(input) do
    hex =
      :crypto.hash(:md5, input)
      |> :binary.bin_to_list()

    %Identicon.Image{hex: hex}
  end
end
