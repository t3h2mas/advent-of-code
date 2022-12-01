package xyz.t3h2mas.advent.three;

import xyz.t3h2mas.advent.three.helpers.SetHelper;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution {
    private static final String inputFilePath = "src/main/java/xyz/t3h2mas/advent/three/input.txt";

    public static void main(String[] args) throws IOException {
        String contents = new String(Files.readAllBytes(Paths.get(inputFilePath)));
        try {
            int closestDistance = wireDistanceClosestToCentralPort(contents);
            System.out.println("Closest distance: " + closestDistance);
        } catch (Exception e) {
            System.err.println(e);
        }

        try {
            int lowestSignal = wireDistanceLowestSignalDelay(contents);
            System.out.println("Lowest signal: " + lowestSignal);
        } catch (Exception e) {
            System.err.println(e);
        }
    }

    public static int wireDistanceClosestToCentralPort(String input) throws Exception {
        List<String>[] wirePaths = SchematicReader.readFrom(input);

        List<String> wirePathOne = wirePaths[0];
        List<String> wirePathTwo = wirePaths[1];

        Point wireOnePosition = new Point(0, 0);
        Point wireTwoPosition = new Point(0, 0);

        Set<Point> wireOneHistory = new HashSet<>();
        Set<Point> wireTwoHistory = new HashSet<>();

        for (int i = 0; i < wirePathOne.size(); i++) {
            final String direction = wirePathOne.get(i).substring(0, 1);
            final int count = Integer.parseInt(wirePathOne.get(i).substring(1));

            for (int j = 0; j < count; j++) {
                switch (direction) {
                    case "R":
                        wireOnePosition = wireOnePosition.moveRight();
                        break;
                    case "L":
                        wireOnePosition = wireOnePosition.moveLeft();
                        break;
                    case "U":
                        wireOnePosition = wireOnePosition.moveUp();
                        break;
                    case "D":
                        wireOnePosition = wireOnePosition.moveDown();
                        break;
                    default:
                        throw new Exception("Unexpected direction: " + direction);
                }

                wireOneHistory.add(wireOnePosition);
            }
        }

        for (int i = 0; i < wirePathTwo.size(); i++) {
            final String direction = wirePathTwo.get(i).substring(0, 1);
            final int count = Integer.parseInt(wirePathTwo.get(i).substring(1));

            for (int j = 0; j < count; j++) {
                switch (direction) {
                    case "R":
                        wireTwoPosition = wireTwoPosition.moveRight();
                        break;
                    case "L":
                        wireTwoPosition = wireTwoPosition.moveLeft();
                        break;
                    case "U":
                        wireTwoPosition = wireTwoPosition.moveUp();
                        break;
                    case "D":
                        wireTwoPosition = wireTwoPosition.moveDown();
                        break;
                    default:
                        throw new Exception("Unexpected direction: " + direction);
                }

                wireTwoHistory.add(wireTwoPosition);
            }
        }


        Set<Point> intersections = SetHelper.intersect(wireOneHistory, wireTwoHistory);

        System.out.println("Intersections found: " + intersections.size());

        int closestDistance = 1000000;
        for (Point intersection : intersections) {
            int distance = ManhattanDistance.calculate(0, intersection.getX(), 0, intersection.getY());
            if (distance < closestDistance) {
                closestDistance = distance;
            }
        }

        return closestDistance;
    }

    public static int wireDistanceLowestSignalDelay(String input) throws Exception {
        List<String>[] wirePaths = SchematicReader.readFrom(input);

        List<String> wirePathOne = wirePaths[0];
        List<String> wirePathTwo = wirePaths[1];

        PointCursor wireOnePosition = new PointCursor(new Point(0, 0));
        PointCursor wireTwoPosition = new PointCursor(new Point(0, 0));

        Set<PointCursor> wireOneHistory = new HashSet<>();
        Set<PointCursor> wireTwoHistory = new HashSet<>();

        for (int i = 0; i < wirePathOne.size(); i++) {
            final String direction = wirePathOne.get(i).substring(0, 1);
            final int count = Integer.parseInt(wirePathOne.get(i).substring(1));

            for (int j = 0; j < count; j++) {
                wireOnePosition = wireOnePosition.moveDirection(direction);
                wireOneHistory.add(wireOnePosition);
            }
        }

        for (int i = 0; i < wirePathTwo.size(); i++) {
            final String direction = wirePathTwo.get(i).substring(0, 1);
            final int count = Integer.parseInt(wirePathTwo.get(i).substring(1));

            for (int j = 0; j < count; j++) {
                wireTwoPosition = wireTwoPosition.moveDirection(direction);
                wireTwoHistory.add(wireTwoPosition);
            }
        }


        Set<PointCursor> intersections = SetHelper.intersect(wireOneHistory, wireTwoHistory);
        System.out.println("Intersections found: " + intersections.size());

        int lowestSignal = 1000000;
        for (PointCursor intersection : intersections) {
            int intersectionSignal = SetHelper.get(wireOneHistory, intersection).getSteps() +
                    SetHelper.get(wireTwoHistory, intersection).getSteps();

            if (intersectionSignal < lowestSignal) {
                lowestSignal = intersectionSignal;
            }
        }

        return lowestSignal;
    }
}
