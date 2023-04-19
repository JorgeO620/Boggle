import { Component } from '@angular/core';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrls: ['./home.component.css']
})
export class HomeComponent {

	jobs: any[];
	news: string[];
	events: string[];

	constructor() {
		this.jobs = [
			{
				title: 'Google recuriter',
				description: ` Hello I am a former senior google sales executive ready to 
				help you in preparing for interview questions and a career here at google.- 
				anonymous user`
			}, {
				title: 'Microsoft recruiter',
				description: `Hello I am a junior marketing coordinator here at Microsoft ready to 
				guide you on an easy application process to a career at Microsoft.`
			}
		];

		this.news = [
			'Your interest, passion, academic goals',
			'Your desired programs, universities, and courses'
		];
		this.events = [
			'Your interest, passion, academic goals',
			'Your desired programs, universities, and courses'
		];
	}

}
