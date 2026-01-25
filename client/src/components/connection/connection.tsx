import type { WebSocketClient } from "../../services/websocket";
import ConnectionForm from "../connection-form/ConnectionForm";

interface ConnectionViewProps {
	websocket: WebSocketClient;
}

export const ConnectionComponent: React.FC<ConnectionViewProps> = ({ websocket }) => {
	if (websocket) {
		console.log("DEBUT FOR WEBSOCKET");
		console.log(websocket)
	}

	return (
		<div className="connection-page">
			<div className="connection-container">
				<ConnectionForm
					initialConnectionId={undefined}
					isLoading={false}
					websocket={websocket}
				/>

			</div>
		</div>
	);
}

export default ConnectionComponent;
