package xyz.t3h2mas.advent.three;

import java.util.Objects;

public class PointCursor {
    private Point point;
    private int steps;

    public PointCursor(Point point) {
        this.point = point;
        this.steps = 0;
    }

    public PointCursor(Point point, int steps) {
        this.point = point;
        this.steps = steps;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        PointCursor that = (PointCursor) o;
        return
                Objects.equals(point, that.point);
    }

    @Override
    public String toString() {
        return "PointCursor{" +
                "point=" + point +
                ", steps=" + steps +
                '}';
    }

    @Override
    public int hashCode() {
        return Objects.hash(point);
    }

    public Point getPoint() {
        return point;
    }

    public int getSteps() {
        return steps;
    }

    public PointCursor moveDirection(String direction) throws Exception {
        switch (direction) {
            case "R":
                return new PointCursor(this.point.moveRight(), this.steps + 1);
            case "L":
                return new PointCursor(this.point.moveLeft(), this.steps + 1);
            case "U":
                return new PointCursor(this.point.moveUp(), this.steps + 1);
            case "D":
                return new PointCursor(this.point.moveDown(), this.steps + 1);
            default:
                throw new Exception("Unexpected direction: " + direction);
        }
    }
}
