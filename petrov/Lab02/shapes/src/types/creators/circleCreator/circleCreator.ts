import { CircleShape } from "../../circle/circleShape";
import { Shape } from "../../shape/shape";
import { ShapeCreator } from "../shapeCreator";

export class CircleCreator implements ShapeCreator {
    private static instance: CircleCreator;

    private constructor() {}

    public static getInstance(): CircleCreator {
        if (!CircleCreator.instance) {
            CircleCreator.instance = new CircleCreator();
        }
        return CircleCreator.instance;
    }

    public createShape(): Shape {
        return new CircleShape(100, 100, 50);
    }
}
