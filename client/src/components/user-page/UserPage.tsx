import type { WebSocketClient } from "../../services/websocket";
import type { IUser } from "../../class/user";

interface UserPageComponentProps {
	websocket?: WebSocketClient;
	user?: IUser
}

export const UserPageComponent: React.FC<UserPageComponentProps> = ({ websocket, user }) => {
	if (!websocket) return (
		<div> no connection </div>
	)

	if (!user) return (
		<div> no user </div>
	)

	const teamNames = {
		1: 'Templiers',
		2: 'Hospitaliers',
		3: 'Scéculiers',
		4: 'Lazarites',
		5: 'Teutoniques'
	};

	const groupNames = {
		1: 'Unité éclaireur',
		2: 'Unit éclaireuse',
		3: 'Unit pionnier',
		4: 'Unit pionnière',
		5: 'Unit animateur'
	};

	return (
		<div className="user-profile-container">
			<div className="connection-card">
				<div className="connection-header">
					<h1 className="connection-title">Profile de Chevalier</h1>
					<p className="connection-subtitle">Vos informations personnlles</p>
				</div>

				{/* User Information Display */}
				<div className="user-info-display">
					<div className="info-card">
						<div className="info-icon">♔</div>
						<div className="info-content">
							<h3>Knight's Name</h3>
							<p id="displayUserName">{user.name}</p>
						</div>
					</div>

					<div className="info-card">
						<div className="info-icon">🆔</div>
						<div className="info-content">
							<h3>Knight's ID</h3>
							<p id="displayUserId">{user.userID}</p>
						</div>
					</div>

					<div className="info-card">
						<div className="info-icon">🏰</div>
						<div className="info-content">
							<h3>Order (Team)</h3>
							<p id="displayUserTeam">
								{teamNames[user.team]}
							</p>
						</div>
					</div>

					<div className="info-card">
						<div className="info-icon">⚜</div>
						<div className="info-content">
							<h3>Unité scout</h3>
							<p id="displayUserGroup">
								{groupNames[user.group]}
							</p>
						</div>
					</div>
				</div>

				{/* Status Indicator 
				<div className="user-status">
					<div className="status-card">
						<div className="status-icon">🛡</div>
						<div className="status-content">
							<h3>Status</h3>
							<p className="status-active">Active Knight</p>
						</div>
					</div>
				</div>

				<div className="knight-heraldry">
					<div className="heraldry-card">
						<div className="heraldry-icon">⚔️</div>
						<div className="heraldry-content">
							<h3>Heraldry Details</h3>
							<p>Registered with honor and valor</p>
						</div>
					</div>
				</div>
				*/}

				{/* Back to Connection Button */}
				<div className="profile-actions">
					<button
						className="action-btn secondary-btn"
					// onClick={onBack}
					>
						<i className="icon-info"></i> Return to Connection
					</button>
				</div>

				<div className="connection-footer">
					<p className="footer-text">
						Your noble status is protected by the ancient codes of chivalry.
					</p>
				</div>
			</div>
		</div>
	);
}

export default UserPageComponent

