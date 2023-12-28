import { ConvexShape } from "../../convex/convexShape";
import { Shape } from "../../shape/shape";
import { ShapeCreator } from "../shapeCreator";

export class ConvexCreator implements ShapeCreator {
    private static instance: ConvexCreator;

    private constructor() {}

    public static getInstance(): ConvexCreator {
        if (!ConvexCreator.instance) {
            ConvexCreator.instance = new ConvexCreator();
        }
        return ConvexCreator.instance;
    }

    public createShape(): Shape {
        return new ConvexShape([{x: 100, y: 100}, {x: 200, y: 100}, {x: 150, y: 200}]);
    }
}