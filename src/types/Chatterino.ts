export interface ChatterinoGlobalBadgesResponse {
	badges: Badge[];
}

export interface Badge {
	tooltip: string;
	image1: string;
	image2: string;
	image3: string;
	users: string[];
}
