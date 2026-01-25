import type { WebSocketClient } from "../../services/websocket";
import ConnectionForm from "../connection-form/ConnectionForm";

interface ConnectionViewProps {
	websocket: WebSocketClient;
}
//
// Decorative components
const KnightThemeDecorations: React.FC = () => (
	<>
		<div className="knight-decor knight-left">♞</div>
		<div className="knight-decor knight-right">♘</div>
	</>
);

const Banner: React.FC = () => (
	<div className="knight-banner">
		<div className="banner-content">
			<div className="banner-title">Registre des ordres</div>
		</div>
	</div>
);

export const ConnectionComponent: React.FC<ConnectionViewProps> = ({ websocket }) => {
	if (websocket) {
		console.log("DEBUT FOR WEBSOCKET");
		console.log(websocket)
	}


	return (
		<div className="connection-page">
			<KnightThemeDecorations />
			<Banner />


			<div className="connection-container">
				<ConnectionForm
					onLogin={undefined}
					initialConnectionId={undefined}
					isLoading={false}
				/>

			</div>
		</div>
	);
}

export default ConnectionComponent;
