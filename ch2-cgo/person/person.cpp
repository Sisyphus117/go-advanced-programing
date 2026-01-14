#include "./person.h"

void Person::Delete() { person_delete(person_handle_t(this)); }

void Person::Set(char *name, int age) {
  person_set(person_handle_t(this), name, age);
}

char *Person::GetName(char *buf, int size) {
  return person_get_name(person_handle_t(this), buf, size);
}

int Person::GetAge() { return person_get_age(person_handle_t(this)); }