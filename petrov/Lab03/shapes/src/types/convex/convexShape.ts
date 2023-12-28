import { Point } from '../point/point';
import { ShapeVisitor } from '../shape/shapeVisitor';
import { IConvexShape } from './convex';

export class ConvexShape implements IConvexShape {
    private points: Point[];
    private type: string = 'convex';

    constructor(points: Point[]) {
        this.points = points;
    }

    public getPoints(): Point[] {
        return this.points;
    }

    public setPoints(points: Point[]): void {
        this.points = points;
    }

    public getType(): string {
        return this.type;
    }

    public draw(canvasContext: CanvasRenderingContext2D): void {
        if (this.points.length < 2) {
            return;
        }
        // the triangle
        canvasContext.beginPath();
        [...this.points, this.points[0]].forEach((point, index) => {
            if (index === 0) {
                canvasContext.moveTo(point.x, point.y);
            } else {
                canvasContext.lineTo(point.x, point.y);
            }
        });
        canvasContext.closePath();

        // the outline
        canvasContext.lineWidth = 10;
        canvasContext.strokeStyle = '#666666';
        canvasContext.stroke();

        // the fill color
        canvasContext.fillStyle = "#FFCC00";
        canvasContext.fill();
    }

    public accept(shapeVisitor: ShapeVisitor): void {
        shapeVisitor.visitConvexShape(this);
    }
}
