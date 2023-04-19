import { Component, OnInit } from '@angular/core';

@Component({
	selector: 'app-profile',
	templateUrl: './profile.component.html',
	styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

	loggedInUser: any;

	constructor() {
		this.loggedInUser = null;
	}

	ngOnInit(): void {
		const user = localStorage.getItem('user');
		this.loggedInUser = user ? JSON.parse(user) : null;
	}
}
