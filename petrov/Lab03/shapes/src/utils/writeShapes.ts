import { Shape } from "../types/shape/shape";
import { decorateShape } from "./decorateShape";

export function exportPerimeterAndArea(shapes: Shape[]) {
    const fileData = shapes.reduce((acc, shape) => {
        const decorator = decorateShape(shape);
        let shapeType = shape.getType();
        if (shapeType === 'convex') {
            shapeType = "triangle";
        }
        acc += `${shapeType.toUpperCase()}: P=${decorator?.getPerimeter()}; S=${decorator?.getArea()}\n`;
        return acc;
        }, "");
    const blob = new Blob([fileData], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.download = "output.txt";
    link.href = url;
    link.click();
  }