import { RectangleShape } from "../../rectangle/rectangleShape";
import { IConcreteDecorator } from "../concreteDecorator";
import { RectangleDecorator } from "./rectangleDecorator";

export class ConcreteRectangleDecorator extends RectangleDecorator implements IConcreteDecorator {
    constructor(rectangle: RectangleShape) {
       super(rectangle);
    }

    public getPerimeter(): number {
        const width = this.rectangle.getWidth();
        const height = this.rectangle.getHeight();
        return 2 * (width + height);
    }

    public getArea(): number {
        const width = this.rectangle.getWidth();
        const height = this.rectangle.getHeight();
        return width * height;
    }
}
