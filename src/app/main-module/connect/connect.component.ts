import { Component, TemplateRef } from '@angular/core';
import { BsModalService, BsModalRef } from 'ngx-bootstrap/modal';

@Component({
	selector: 'app-connect',
	templateUrl: './connect.component.html',
	styleUrls: ['./connect.component.css']
})
export class ConnectComponent {

	message: string;
	recuriters: any[];
	selectedRecuriter: any;

	modalRef?: BsModalRef<any>;

	constructor(private modalService: BsModalService) {

		this.message = '';
		this.selectedRecuriter = null;

		this.recuriters = [{
			id: 1,
			name: 'Katlyn Perkins',
			role: 'Google recruiter',
			image: 'https://tse2.mm.bing.net/th?id=OIP.HwTtqc10QA1N0_DX5ruW7QHaEJ&pid=Api&P=0',
			isInviteSended: true
		}, {
			id: 2,
			name: 'Max Matherson',
			role: 'Microsoft recruiter',
			image: 'https://tse1.mm.bing.net/th?id=OIP.bFNFNtbodjO1HjAk4Ek-1QHaHa&pid=Api&P=0',
			isInviteSended: false 
		}, {
			id: 3,
			name: 'Steven King',
			role: 'Amazon tech recruiter',
			image: 'https://tse2.mm.bing.net/th?id=OIP.2SvJkiywOojR2PuqiNwN6AHaEK&pid=Api&P=0',
			isInviteSended: false 
		}, {
			id: 4,
			name: 'Walter Patterson',
			role: 'Apple recruiter',
			image: 'https://tse3.mm.bing.net/th?id=OIP.iDetnmIJeB_t79HZFWpIkAHaE8&pid=Api&P=0',
			isInviteSended: false 
		}, {
			id: 5,
			name: 'Phillip Anderson',
			role: 'J.P Morgan Recruiter',
			image: 'https://tse3.mm.bing.net/th?id=OIP.eKTBlb4IZ5UuFavcpylTNgAAAA&pid=Api&P=0',
			isInviteSended: false 
		}];
	}

	onConnect(recuriter: any) {
		recuriter['loading'] = true;

		setTimeout(() => {
			recuriter.isInviteSended = true;
			recuriter.loading = false;
		}, 1000);
	}

	openModal(recuriter: any, template: TemplateRef<any>) {
		this.selectedRecuriter = recuriter
		this.modalRef = this.modalService.show(template);
	}

	onSendMessage() {
		this.selectedRecuriter['message_sending'] = true;

		setTimeout(() => {
			const recruiter = this.recuriters.find(r => r.id === this.selectedRecuriter.id);
			if (recruiter) {
				this.modalRef?.hide();
				recruiter.isInviteSended = true;
				this.selectedRecuriter = null;
			}
		}, 1000);
	}
}
