import { Point } from '../types/point/point';
import { Shape } from '../types/shape/shape';
import { CircleCreator } from '../types/creators/circleCreator/circleCreator';
import { RectangleCreator } from '../types/creators/rectangleCreator/rectangleCreator';
import { ConvexCreator } from '../types/creators/convexCreator/convexCreator';
import { isConvexShape } from '../types/convex/guard';
import { isRectangleShape } from '../types/rectangle/guard';
import { isCircleShape } from '../types/circle/guard';


export async function readShapesFromFile(file: File): Promise<Shape[]> {
    const fileData = new FileReader();
    const fileContentPromise = new Promise<string>((resolve) => {
        fileData.onloadend = (e: ProgressEvent<FileReader>) => {
            if (e.target) {
                resolve(e.target.result?.toString() as string);
            }
        };
        fileData.readAsText(file);
    })

    const shapes: Shape[] = [];

    const fileContent = await fileContentPromise;
    const lines = fileContent.split('\n');

    const circleCreator = CircleCreator.getInstance();
    const rectangleCreator = RectangleCreator.getInstance();
    const convexCreator = ConvexCreator.getInstance();

    for (const line of lines) {
        const shapeData = line.trim().split(':');
        const shapeType = shapeData[0].trim();
        const shapeParams = shapeData[1].trim();

        if (shapeType === 'TRIANGLE') {
            const points: Point[] = shapeParams.split(';').map((point) => {
                const [x, y] = point.split('=')[1].split(',');
                return { x: Number(x), y: Number(y) };
            });
            const convexShape = convexCreator.createShape();
            if (isConvexShape(convexShape)) {
                convexShape.setPoints([points[0], points[1], points[2]]);
                shapes.push(convexShape);
            }
        } else if (shapeType === 'RECTANGLE') {
            const [p1, p2] = shapeParams.split(';').map((point) => {
                const [x, y] = point.split('=')[1].split(',');
                return { x: Number(x), y: Number(y) };
            });
            const rectangleShape = rectangleCreator.createShape();
            if (isRectangleShape(rectangleShape)) {
                rectangleShape.setX(p1.x);
                rectangleShape.setY(p1.y);
                rectangleShape.setWidth(p2.x - p1.x);
                rectangleShape.setHeight(p2.y - p1.y);
            }
            shapes.push(rectangleShape);
        } else if (shapeType === 'CIRCLE') {
            let center: Point = {x: 0, y: 0};
            let radius: number = 0;
            shapeParams.split(';').forEach((param) => {

                const [key, value] = param.split('=');
                if (key.trim() === 'C') {
                    const [x, y] = value.split(',');
                    center = { x: Number(x), y: Number(y) };
                } else if (key.trim() === 'R') {
                    radius = Number(value);
                }
            });
            const circleShape = circleCreator.createShape();
            if (isCircleShape(circleShape)) {
                circleShape.setRadius(radius);
                circleShape.setX(center.x);
                circleShape.setY(center.y);
            }
            shapes.push(circleShape);
        }
    }

    return shapes;
}
