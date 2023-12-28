import { CircleShape } from "../circle/circleShape";
import { ConvexShape } from "../convex/convexShape";
import { RectangleShape } from "../rectangle/rectangleShape";
import { ShapeVisitor } from "../shape/shapeVisitor";

export class PerimeterAndAreaVisitor implements ShapeVisitor {
    private writeData: string = "";

    private writePerimeterAndArea(shapeType: string, perimeter: number, area: number): void {
        if (shapeType === 'convex') {
            shapeType = "triangle";
        }
        this.writeData += `${shapeType.toUpperCase()}: P=${perimeter}; S=${area}\n`;
    }

    private calculateCirclePerimeter(circle: CircleShape): number {
        return 2 * Math.PI * circle.getRadius();
    }

    private calculateCircleArea(circle: CircleShape): number {
        return Math.PI * Math.pow(circle.getRadius(), 2);
    }

    public visitCircleShape(circleShape: CircleShape): void {
        const perimeter = this.calculateCirclePerimeter(circleShape);
        const area = this.calculateCircleArea(circleShape);
        this.writePerimeterAndArea('circle', perimeter, area);
    }

    private calculateRectanglePerimeter(rectangleShape: RectangleShape): number {
        const width = rectangleShape.getWidth();
        const height = rectangleShape.getHeight();
        return 2 * (width + height);
    }

    private calculateRectangleArea(rectangleShape: RectangleShape): number {
        const width = rectangleShape.getWidth();
        const height = rectangleShape.getHeight();
        return width * height;
    }

    public visitRectangleShape(rectangleShape: RectangleShape): void {
        const perimeter = this.calculateRectanglePerimeter(rectangleShape);
        const area = this.calculateRectangleArea(rectangleShape);
        this.writePerimeterAndArea('rectangle', perimeter, area);
    }

    private calculateConvexPerimeter(convexShape: ConvexShape): number {
        const points = convexShape.getPoints();
        let perimeter = 0;

        for (let i = 0; i < points.length; i++) {
            const point1 = points[i];
            const point2 = points[(i + 1) % points.length]; // Wrap around to the first point for the last side
            perimeter += Math.sqrt(Math.pow(point2.x - point1.x, 2) + Math.pow(point2.y - point1.y, 2));
        }

        return perimeter;
    }

    private calculateConvexArea(convexShape: ConvexShape): number {
        const points = convexShape.getPoints();
        let area = 0;

        for (let i = 0; i < points.length - 1; i++) {
            const point1 = points[i];
            const point2 = points[i + 1];
            area += (point1.x * point2.y) - (point2.x * point1.y);
        }

        return Math.abs(area) / 2;
    }

    public visitConvexShape(convexShape: ConvexShape): void {
        const perimeter = this.calculateConvexPerimeter(convexShape);
        const area = this.calculateConvexArea(convexShape);
        this.writePerimeterAndArea('convex', perimeter, area);
    }

    public downloadData(fileName: string): void {
        const blob = new Blob([this.writeData], { type: "text/plain" });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.download = fileName;
        link.href = url;
        link.click();
    }

    public getWriteData(): string {
        return this.writeData;
    }

    public setWriteData(writeData: string): void {
        this.writeData = writeData;
    }
}