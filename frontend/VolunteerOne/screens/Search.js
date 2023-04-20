import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import NotificationItem from '../components/NotificationItem';
import notifications from '../constants/notifications';

/** ============================== Search Page ============================== **/ 
const { width } = Dimensions.get('screen');

class SearchPage extends React.Component {
  renderNotifications = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
        
        {/* map all notifications using loop */}
        <Block flex>
          {/* {notifications.map(notification => (
           <NotificationItem item={notification} horizontal />
          ))} */}
        </Block>
    
      </ScrollView>
    )
  }

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderNotifications()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default SearchPage;
