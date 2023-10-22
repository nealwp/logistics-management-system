import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';

export interface Item {
    id: string
    name: string
    unitPrice: number
}

const fakeItems: Item[] = [
    {id: "289819873", name: "bolt", unitPrice: 1.23},
    {id: "282873936", name: "screwdriver", unitPrice: 6.12},
    {id: "228379874", name: "broom handle", unitPrice: 5.34},
    {id: "293849872", name: "mop bucker", unitPrice: 12.99},
    {id: "293849239", name: "cleaning wipes", unitPrice: 4.17},
    {id: "219872364", name: "pipe cleaners", unitPrice: 2.73},
]

@Injectable({
  providedIn: 'root'
})
export class ItemService {

    constructor(private http: HttpClient) {}

    getItems(): Observable<Item[]> {
        return of(fakeItems)
    }

    getItemById(id: string): Observable<Item> {
        const item = fakeItems.find(entry => entry.id === id) 
        if (item) {
            return of(item)
        } else {
            return of({id: "", name: "", unitPrice: 0.00})
        }
    }
}


