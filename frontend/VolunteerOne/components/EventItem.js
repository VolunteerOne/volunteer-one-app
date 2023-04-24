import React from "react";
import PropTypes from "prop-types";
import { StyleSheet, TouchableWithoutFeedback, Alert } from "react-native";
import { Block, Text, theme } from "galio-framework";
import { Button } from "../components";
import { Avatar } from "@react-native-material/core";
import { useNavigation } from "@react-navigation/native";

/*
Event Item component receives the details of the event such as organization, title, and event id.
The event items are displayed in your bookmark page which is accessible from your user profile. 
If you click on an event item, it will take you to its corresponding event post. 
You also have the ability to remove events from your bookmark by clicking on the remove button 

props
data - Type (obj):                  {organization, name, id}
handleRemoveItem - Type (func):     handles removal of the item
*/
const EventItem = (props) => {
  //receive props
  const handleRemoveItem = props.handleRemoveItem;
  const eventInfo = props.data;
  //navigation used to navigate to event details page if event item is clicked on
  const navigation = useNavigation();
  //styling
  const imageStyles = [styles.horizontalImage];
  const cardContainer = [styles.card, styles.shadow];

  /*
    When user taps on remove, they are asked to give confirmation. If they select
    yes, then the item is removed from the bookmark. If no is selected, the action 
    is cancled 
    
    parameters 
    id - Type (int):    event id that was selected for removal 
  */
  const confirmationAlert = (id) => {
    Alert.alert("Are you sure?", "", [
      {
        text: "Cancel",
        onPress: () => console.log("Cancel Pressed"),
        style: "cancel",
      },
      { text: "OK", onPress: () => handleRemoveItem(id) },
    ]);
  };
  return (
    <TouchableWithoutFeedback
      onPress={() =>
        navigation.navigate("ViewEvent", { eventID: eventInfo["id"] })
      }
    >
      <Block row={true} card style={cardContainer}>
        <Block>
          <Avatar
            style={imageStyles}
            image={{ uri: eventInfo["image"] }}
            color="white"
          />
        </Block>

        <Block flex>
          <Text size={12} style={styles.cardTitle} bold>
            {eventInfo["organization"]}
          </Text>
          <Text size={12}>{eventInfo["name"]}</Text>
        </Block>

        <Block row={true} style={styles.cardDescription}>
          <Button
            small
            style={{ backgroundColor: "grey" }}
            onPress={() => confirmationAlert(eventInfo["id"])}
          >
            Remove
          </Button>
        </Block>
      </Block>
    </TouchableWithoutFeedback>
  );
};

EventItem.propTypes = {
  data: PropTypes.object,
  handleRemoveItem: PropTypes.func,
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: theme.COLORS.WHITE,
    marginVertical: theme.SIZES.BASE / 2,
    minHeight: 96,
    justifyContent: "flex-start",
    alignItems: "center",
  },
  cardTitle: {},
  cardDescription: {
    marginRight: 20,
    borderRadius: 3,
    elevation: 1,
    overflow: "hidden",
  },
  image: {},
  horizontalImage: {
    height: 50,
    width: 50,
    borderRadius: 62,
    margin: 20,
  },
  horizontalStyles: {
    borderTopRightRadius: 0,
    borderBottomRightRadius: 0,
  },
  verticalStyles: {
    borderBottomRightRadius: 0,
    borderBottomLeftRadius: 0,
  },
  fullImage: {
    height: 215,
  },
  shadow: {
    shadowColor: theme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.1,
    elevation: 2,
  },
});

export default EventItem;
