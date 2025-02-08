#include <iostream>
#include <vector>
#include <memory>
#include <algorithm>

class Item {
public:
    Item(int id, std::string name) : id_(id), name_(name) {}
    int getId() const { return id_; }
    std::string getName() const { return name_; }

private:
    int id_;
    std::string name_;
};

class ItemManager {
public:
    void addItem(std::shared_ptr<Item> item) {
        items_.push_back(item);
    }

    std::shared_ptr<Item> findItemById(int id) {
        auto it = std::find_if(items_.begin(), items_.end(),
            [id](const std::shared_ptr<Item>& item) { return item->getId() == id; });
        if (it != items_.end()) {
            return *it;
        }
        return nullptr;
    }

    void printItems() const {
        for (const auto& item : items_) {
            std::cout << "Item ID: " << item->getId() << ", Name: " << item->getName() << std::endl;
        }
    }

private:
    std::vector<std::shared_ptr<Item>> items_;
};

int main() {
    ItemManager manager;

    manager.addItem(std::make_shared<Item>(1, "Item1"));
    manager.addItem(std::make_shared<Item>(2, "Item2"));
    manager.addItem(std::make_shared<Item>(3, "Item3"));

    manager.printItems();

    auto item = manager.findItemById(2);
    if (item) {
        std::cout << "Found Item: " << item->getName() << std::endl;
    } else {
        std::cout << "Item not found" << std::endl;
    }

    return 0;
}
