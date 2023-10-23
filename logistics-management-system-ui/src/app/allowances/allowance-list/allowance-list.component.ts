import { Component } from '@angular/core';
import { Allowance, AllowanceService } from '../allowance.service';

@Component({
  selector: 'app-allowance-list',
  templateUrl: './allowance-list.component.html',
  styleUrls: ['./allowance-list.component.scss']
})
export class AllowanceListComponent {
    allowances: Allowance[] = [];

    constructor(private allowanceService: AllowanceService) {};

    ngOnInit(): void {
        this.allowanceService.getAllowances().subscribe(data => {
            this.allowances = data;
        })
    }
}
