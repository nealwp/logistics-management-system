import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

export interface Allowance {
    itemId: string
    level: number
}

const fakeAllowances: Allowance[] = [
    {itemId: "289819873", level: 23},
    {itemId: "282873936", level: 12},
    {itemId: "228379874", level: 34},
    {itemId: "293849872", level: 99},
    {itemId: "293849239", level: 17},
    {itemId: "219872364", level: 73},
]

@Injectable({
  providedIn: 'root'
})
export class AllowanceService {

    constructor() { }

    getAllowances(): Observable<Allowance[]> {
        return of(fakeAllowances)
    }

    getAllowanceByItemId(itemId: string): Observable<number> {
        const allowance = fakeAllowances.find(entry => entry.itemId === itemId) 
        if (allowance) {
            return of(allowance.level)
        } else {
            return of(0)
        }
    }
}
