import { StyleSheet } from "react-native";
import { Avatar } from "@rneui/themed";
import { Block, Text } from "galio-framework";

/*this function gets the difference between the current date and the date the event was posted. It then formats
the time displayed accordingly */
function getDateDiff(date1, date2) {
  let timeDisplayed = "";
  // Get the difference in milliseconds
  const diffInMs = Math.abs(date1 - date2);

  // Convert milliseconds to days, hours, and minutes
  const days = Math.floor(diffInMs / (1000 * 60 * 60 * 24));
  const hours = Math.floor(
    (diffInMs % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
  );
  const minutes = Math.floor((diffInMs % (1000 * 60 * 60)) / (1000 * 60));

  //get date values of the post date//
  var day = date2.getUTCDate();
  var month = date2.toLocaleString("default", { month: "long" });
  var year = date2.getUTCFullYear();

  //if days greater than 7 and the year the event was posted was a previous year, display month, day, and year
  if (days > 7 && date1.getUTCFullYear() > year) {
    timeDisplayed = month + " " + day + ", " + year;
    //if days greater than 7, and the year the event was posted is this year, display month and day
  } else if (days > 7 && date1.getUTCFullYear() == year) {
    timeDisplayed = month + " " + day;
    //if days between 1-7, display days
  } else if (days > 0) {
    const formatDays = days == 1 ? "day" : "days";
    timeDisplayed = days + " " + formatDays + " ago";
    //if hours greater than 0, display hours
  } else if (hours > 0) {
    const formatHours = hours == 1 ? "hour" : "hours";
    timeDisplayed = hours + " " + formatHours + " ago";
    //if minutes greater than 0, display minutes
  } else if (minutes > 0) {
    const formatMinutes = minutes == 1 ? "minute" : "minutes";
    timeDisplayed = minutes + " " + formatMinutes + " ago";
    //if post was made less than a minute ago, say posted just now
  } else {
    timeDisplayed = "Just Now";
  }
  // Return the result as an object
  return timeDisplayed;
}

/*
Description:
  This component returns the header of a card.
Props received:
  organization - String
  timePosted - Date object 
*/
const CardHeader = ({ organization, timePosted }) => {
  //getting time to display
  const timeDisplayed = getDateDiff(new Date(), timePosted);
  return (
    <Block style={styles.header}>
      <Avatar
        size={62}
        rounded
        containerStyle={[styles.avatarStyle, styles.shadowProp]}
      />
      <Block style={styles.headerTitle}>
        <Text style={styles.titleText}>{organization}</Text>
        <Text style={styles.timeText}>{timeDisplayed}</Text>
      </Block>
    </Block>
  );
};

const styles = StyleSheet.create({
  titleText: {
    fontSize: 20,
    fontWeight: "bold",
    color: "#32325D",
  },
  timeText: {
    color: "#32325D",
  },
  header: {
    flexDirection: "row",
    justifyContent: "flex-start",
    marginBottom: 10,
  },
  headerTitle: {
    marginLeft: 10,
  },
  avatarStyle: {
    backgroundColor: "gray",
  },
  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default CardHeader;
