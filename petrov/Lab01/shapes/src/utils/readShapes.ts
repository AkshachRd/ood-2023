import { Point } from '../types/point/point';
import { Shape } from '../types/shape/shape';
import { ConvexShape } from '../types/convex/convexShape';
import { RectangleShape } from '../types/rectangle/rectangleShape';
import { CircleShape } from '../types/circle/circleShape';


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

    for (const line of lines) {
        const shapeData = line.trim().split(':');
        const shapeType = shapeData[0].trim();
        const shapeParams = shapeData[1].trim();

        if (shapeType === 'TRIANGLE') {
            const points: Point[] = shapeParams.split(';').map((point) => {
                const [x, y] = point.split('=')[1].split(',');
                return { x: Number(x), y: Number(y) };
            });
            shapes.push(new ConvexShape([points[0], points[1], points[2]]));
        } else if (shapeType === 'RECTANGLE') {
            const [p1, p2] = shapeParams.split(';').map((point) => {
                const [x, y] = point.split('=')[1].split(',');
                return { x: Number(x), y: Number(y) };
            });
            shapes.push(new RectangleShape(p1.x, p1.y, p2.x - p1.x, p2.y - p1.y));
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

            shapes.push(new CircleShape(center.x, center.y, radius));
        }
    }

    return shapes;
}
