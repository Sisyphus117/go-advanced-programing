extern "C" {
#include "./person_capi.h"
}
struct Person {
public:
  static Person *New(const char *name, int age) {
    return (Person *)person_new((char *)name, age);
  }
  void Delete();

  void Set(char *name, int age);
  char *GetName(char *buf, int size);
  int GetAge();
};
