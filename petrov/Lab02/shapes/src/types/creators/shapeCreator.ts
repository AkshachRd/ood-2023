import { Shape } from "../shape/shape";

export abstract class ShapeCreator {
    abstract createShape(): Shape
}