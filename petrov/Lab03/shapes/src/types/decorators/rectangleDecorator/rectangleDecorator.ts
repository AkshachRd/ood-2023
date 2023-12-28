import { IRectangleShape } from "../../rectangle/rectangle";
import { RectangleShape } from "../../rectangle/rectangleShape";
import { ShapeVisitor } from "../../shape/shapeVisitor";

export class RectangleDecorator implements IRectangleShape {
    protected rectangle: RectangleShape;

    constructor(rectangle: RectangleShape) {
        this.rectangle = rectangle;
    }

    public getWidth(): number {
        return this.rectangle.getWidth();
    }

    public setWidth(width: number): void {
        this.rectangle.setWidth(width);
    }

    public getHeight(): number {
        return this.rectangle.getHeight();
    }

    public setHeight(height: number): void {
        this.rectangle.setHeight(height);
    }

    public getY(): number {
        return this.rectangle.getY();
    }

    public setY(y: number): void {
        this.rectangle.setY(y);
    }

    public getX(): number {
        return this.rectangle.getX();
    }

    public setX(x: number): void {
        this.rectangle.setX(x);
    }

    public getType(): string {
        return this.rectangle.getType();
    }

    public draw(canvasContext: CanvasRenderingContext2D): void {
        this.rectangle.draw(canvasContext);
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

    public accept(shapeVisitor: ShapeVisitor): void {
        this.rectangle.accept(shapeVisitor);
    }
}
