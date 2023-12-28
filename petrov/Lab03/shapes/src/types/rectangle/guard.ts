import { Shape } from "../shape/shape"
import { RectangleShape } from "./rectangleShape"

export const isRectangleShape = (shape: Shape): shape is RectangleShape => {
    return shape.getType() === "rectangle" &&
        (shape as RectangleShape).getHeight !== undefined &&
        (shape as RectangleShape).setHeight !== undefined &&
        (shape as RectangleShape).getWidth !== undefined &&
        (shape as RectangleShape).setWidth !== undefined &&
        (shape as RectangleShape).getX !== undefined && (shape as RectangleShape).setX !== undefined &&
        (shape as RectangleShape).getY !== undefined && (shape as RectangleShape).setY !== undefined
}