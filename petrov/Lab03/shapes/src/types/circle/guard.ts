import { Shape } from "../shape/shape";
import { CircleShape } from "./circleShape";

export const isCircleShape = (shape: Shape): shape is CircleShape => {
    return shape.getType() === "circle" &&
        (shape as CircleShape).getRadius !== undefined &&
        (shape as CircleShape).getX !== undefined && (shape as CircleShape).getY !== undefined &&
        (shape as CircleShape).setRadius !== undefined &&
        (shape as CircleShape).setX !== undefined && (shape as CircleShape).setY !== undefined
}