package kata

import "testing"

func TestBob27Male(t *testing.T) {
	//Arrange
	dm := NewDinglemouse().SetName("Bob").SetAge(27).SetSex('M')
	expected := "Hello. My name is Bob. I am 27. I am male."

	//Act
	actual := dm.Hello()

	//Assert
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func Test27MaleBob(t *testing.T) {
	//Arrange
	dm := NewDinglemouse().SetAge(27).SetSex('M').SetName("Bob")
	expected := "Hello. I am 27. I am male. My name is Bob."

	//Act
	actual := dm.Hello()

	//Assert
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestAliceFemale(t *testing.T) {
	//Arrange
	dm := NewDinglemouse().SetName("Alice").SetSex('F')
	expected := "Hello. My name is Alice. I am female."

	//Act
	actual := dm.Hello()

	//Assert
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestBatman(t *testing.T) {
	//Arrange
	dm := NewDinglemouse().SetName("Batman")
	expected := "Hello. My name is Batman."

	//Act
	actual := dm.Hello()

	//Assert
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
