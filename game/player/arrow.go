components {
  id: "arrow"
  component: "/game/player/arrow.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"arrow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/resources/sprites/atlas.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.07
    y: 0.07
  }
}
