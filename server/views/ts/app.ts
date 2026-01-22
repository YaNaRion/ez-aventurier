// Import modules
import { ControllerService } from './service/controller.js';

interface GlobalState {
	controllerService: ControllerService;
}

const global_state: GlobalState = {
	controllerService: new ControllerService(),
};


// Initialize application
function init() {
	console.log('Application initializing...');
	// Initialise top bar
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
	document.addEventListener('DOMContentLoaded', init);
} else {
	init();
}
