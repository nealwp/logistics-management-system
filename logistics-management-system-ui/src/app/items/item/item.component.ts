import { Component, OnInit } from '@angular/core';
import { Item, ItemService } from '../item.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-item',
  templateUrl: './item.component.html',
  styleUrls: ['./item.component.scss']
})
export class ItemComponent implements OnInit {
    item: Item = {id: "", name: "", unitPrice: 0.00};

    constructor(private route: ActivatedRoute, private itemService: ItemService) {}

    ngOnInit(): void {
        const itemId = this.route.snapshot.paramMap.get('id');
        if (itemId) {
            this.itemService.getItemById(itemId)
                .subscribe(item => { this.item = item })
        }
    }   
}
