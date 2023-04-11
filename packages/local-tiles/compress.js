import sharp from "sharp";

//22.7 ko
const path = "tiles/13/5011/4382.png";
sharp(path).jpeg({ mozjpeg: true }).toFile("out.jpeg");
