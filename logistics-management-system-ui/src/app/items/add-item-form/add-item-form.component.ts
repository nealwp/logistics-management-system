import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Item, ItemService } from '../item.service';

@Component({
  selector: 'app-add-item-form',
  templateUrl: './add-item-form.component.html',
  styleUrls: ['./add-item-form.component.scss']
})
export class AddItemFormComponent {

  item: Item = <Item>{}

  constructor(private router: Router, private itemService: ItemService){};

  public onSubmit() {
      this.itemService.createItem(this.item)
      this.router.navigate(['/item', this.item.id]);
  }

  public onCancel() {
      this.router.navigate(['/']);
  }
}
