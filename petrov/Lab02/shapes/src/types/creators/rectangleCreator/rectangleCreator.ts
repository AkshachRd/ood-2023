import { RectangleShape } from "../../rectangle/rectangleShape";
import { Shape } from "../../shape/shape";
import { ShapeCreator } from "../shapeCreator";

export class RectangleCreator implements ShapeCreator {
    private static instance: RectangleCreator;

    private constructor() {}

    public static getInstance(): RectangleCreator {
        if (!RectangleCreator.instance) {
            RectangleCreator.instance = new RectangleCreator();
        }
        return RectangleCreator.instance;
    }

    public createShape(): Shape {
        return new RectangleShape(100, 100, 100, 200);
    }
}
