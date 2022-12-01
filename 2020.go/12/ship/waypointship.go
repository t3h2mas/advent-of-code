package ship

type WaypointShip struct {
	waypoint *Waypoint
	ship     *Ship
}

func NewWaypointShip() *WaypointShip {
	return &WaypointShip{
		NewWaypoint(),
		NewShip(),
	}
}

func (ws *WaypointShip) Act(action rune, units int) {
	switch action {
	case 'F':
		for i := 0; i < units; i++ {
			ws.ship.move(North, ws.waypoint.position.Y())
			ws.ship.move(West, ws.waypoint.position.X())
		}
	default:
		ws.waypoint.Act(action, units)
	}
}

func (ws *WaypointShip) ShipPosition() Position {
	return ws.ship.Position()
}
